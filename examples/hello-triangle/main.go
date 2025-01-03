// Copyright 2024 Roi Martin.

// Hello-triangle is a simple program that shows how to use imgui with
// the GLFW and OpenGL backends.
package main

import (
	"log"
	"runtime"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/jroimartin/imgui"
	igglfw "github.com/jroimartin/imgui/backend/glfw"
	igopengl "github.com/jroimartin/imgui/backend/opengl"
)

const (
	initialWidth       = 800
	initialHeight      = 600
	vertexShaderSource = `
		#version 330 core
		layout (location = 0) in vec3 aPos;

		void main()
		{
			gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
		}
	`
	fragmentShaderSource = `
		#version 330 core
		out vec4 fragColor;

		uniform vec4 uColor;

		void main()
		{
			fragColor = uColor;
		}
	`
)

var vertices = []float32{
	-0.5, -0.5, 0.0, 0.5, -0.5, 0.0, 0.0, 0.5, 0.0,
}

func main() {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		log.Fatalf("failed to initialize GLFW: %v", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	win, err := glfw.CreateWindow(initialWidth, initialHeight, "Hello Triangle", nil, nil)
	if err != nil {
		log.Fatalf("failed to create GLFW window: %v", err)
	}
	win.MakeContextCurrent()

	win.SetFramebufferSizeCallback(framebufferSizeCB)

	if err := gl.Init(); err != nil {
		log.Fatalf("failed to initialize OpenGL")
	}

	gl.Enable(gl.DEBUG_OUTPUT)
	gl.DebugMessageCallback(glDebugMessageCB, nil)

	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	cstrs, free := gl.Strs(vertexShaderSource + "\x00")
	gl.ShaderSource(vertexShader, 1, cstrs, nil)
	free()
	gl.CompileShader(vertexShader)

	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	cstrs, free = gl.Strs(fragmentShaderSource + "\x00")
	gl.ShaderSource(fragmentShader, 1, cstrs, nil)
	free()
	gl.CompileShader(fragmentShader)

	shaderProgram := gl.CreateProgram()
	defer gl.DeleteProgram(shaderProgram)

	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)
	gl.LinkProgram(shaderProgram)
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	defer gl.DeleteVertexArrays(1, &vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	defer gl.DeleteBuffers(1, &vbo)

	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	igCtx := imgui.CreateContext(nil)
	defer igCtx.Destroy()

	igIO := imgui.GetIO()
	igIO.SetConfigFlags(igIO.ConfigFlags() | imgui.ConfigFlagsNavEnableKeyboard | imgui.ConfigFlagsDockingEnable)
	igIO.SetIniFilename("")
	igIO.SetLogFilename("")

	if !igglfw.InitForOpenGL(win.Handle(), true) {
		log.Fatal("failed to initialize ImGui GLFW backend")
	}
	defer igglfw.Shutdown()

	if !igopengl.Init("#version 330 core") {
		log.Fatal("failed to initialize ImGui OpenGL backend")
	}
	defer igopengl.Shutdown()

	uniformLocation := gl.GetUniformLocation(shaderProgram, gl.Str("uColor\x00"))

	winIsOpen := true
	triangleColor := imgui.Color4{R: 1.0, G: 0.5, B: 0.2, A: 1.0}

	for !win.ShouldClose() {
		glfw.PollEvents()

		igopengl.NewFrame()
		igglfw.NewFrame()
		imgui.NewFrame()

		if winIsOpen {
			vp := imgui.GetMainViewport()
			workpos := vp.GetWorkpos()
			imgui.SetNextWindowPos(imgui.Vec2{X: workpos.X + 10.0, Y: workpos.Y + 10.0}, 0, imgui.Vec2{})
			if imgui.Begin("Configuration", &winIsOpen, imgui.WindowFlagsAlwaysAutoresize) {
				imgui.ColorEdit4("Triangle Color", &triangleColor, imgui.ColorEditFlagsNoInputs)
			}
			imgui.End()
		}

		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.UseProgram(shaderProgram)
		gl.Uniform4f(uniformLocation, triangleColor.R, triangleColor.G, triangleColor.B, triangleColor.A)
		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		imgui.Render()
		igopengl.RenderDrawData(imgui.GetDrawData())

		win.SwapBuffers()
	}
}

func framebufferSizeCB(_win *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}

func glDebugMessageCB(source uint32, typ uint32, id uint32, severity uint32, _length int32, message string, _userParam unsafe.Pointer) {
	log.Printf("GL debug: %v: source=%v type=%v id=%v severity=%v", message, source, typ, id, severity)
}
