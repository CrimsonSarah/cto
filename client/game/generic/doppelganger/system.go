package doppelganger

// import "github.com/CrimsonSarah/cto/client/engine"
// 
// // May assume any kind. Used for testing.
// 
// // Logic system.
// 
// type LogicDoppelganger struct {
// 	engine.SystemBase
// 
// 	kind *engine.NodeKind
// 	tick func(ctx engine.FrameContext, node any)
// }
// 
// func MakeLogicDoppelganger(kind *engine.NodeKind) LogicDoppelganger {
// 	defaultTick := func(ctx engine.FrameContext, node any) {}
// 
// 	return LogicDoppelganger{
// 		kind: kind,
// 		tick: defaultTick,
// 	}
// }
// 
// func (d *LogicDoppelganger) SetTick(tick func(ctx engine.FrameContext, node any)) {
// 	d.tick = tick
// }
// 
// func (d LogicDoppelganger) NodeKind() *engine.NodeKind {
// 	return d.kind
// }
// 
// func (d LogicDoppelganger) Init(engine.InitContext) {}
// func (d LogicDoppelganger) Destroy()                {}
// 
// func (d LogicDoppelganger) Tick(ctx engine.FrameContext, node any) {
// 	d.tick(ctx, node)
// }
// 
// // Render system.
// 
// type RenderDoppelganger struct {
// 	engine.RenderSystemBase
// 
// 	kind   *engine.NodeKind
// 	render func(ctx engine.RenderContext, node any)
// }
// 
// func MakeRenderDoppelganger(kind *engine.NodeKind) RenderDoppelganger {
// 	defaultRender := func(ctx engine.RenderContext, node any) {}
// 
// 	return RenderDoppelganger{
// 		kind:   kind,
// 		render: defaultRender,
// 	}
// }
// 
// func (d *RenderDoppelganger) SetRender(render func(ctx engine.RenderContext, node any)) {
// 	d.render = render
// }
// 
// func (d RenderDoppelganger) NodeKind() *engine.NodeKind {
// 	return d.kind
// }
// 
// func (d RenderDoppelganger) Init(engine.InitContext) {}
// func (d RenderDoppelganger) Destroy()                {}
// 
// func (d RenderDoppelganger) Render(ctx engine.RenderContext, node any) {
// 	d.render(ctx, node)
// }
