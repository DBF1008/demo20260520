package service

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)


type GfWatcher struct {
	ctx      context.Context
	cancel   context.CancelFunc
	callback func(string)
	channel  string
}


func NewGfWatcher(channelName string) (*GfWatcher, error) {
	ctx, cancel := context.WithCancel(context.Background())

	w := &GfWatcher{
		ctx:     ctx,
		cancel:  cancel,
		channel: channelName,
	}


	go w.startSubscribe()

	return w, nil
}


func (w *GfWatcher) SetUpdateCallback(callback func(string)) error {
	w.callback = callback
	return nil
}


func (w *GfWatcher) Update() error {
	_, err := g.Redis().Publish(w.ctx, w.channel, "policy_updated")
	return err
}


func (w *GfWatcher) Close() {

	w.cancel()
}


func (w *GfWatcher) startSubscribe() {
	sub, _, err := g.Redis().Subscribe(w.ctx, w.channel)
	if err != nil {
		g.Log().Error(w.ctx, "Casbin Watcher 订阅失败:", err)
		return
	}

	defer func() {
		_ = sub.Close(w.ctx)
	}()


	for {

		select {
		case <-w.ctx.Done():
			g.Log().Info(w.ctx, "Casbin Watcher 停止监听")
			return
		default:

		}


		msg, err := sub.Receive(w.ctx)
		if err != nil {

			if w.ctx.Err() != nil {
				return
			}

			g.Log().Error(w.ctx, "Casbin Watcher 接收消息错误:", err)

			return
		}


		if w.callback != nil && msg != nil {
			g.Log().Debug(w.ctx, "Casbin Watcher 收到更新通知")
			w.callback("update signal received")
		}
	}
}
