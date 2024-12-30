// Copyright 2024 Roi Martin.

/*
Package imgui provides bindings for the [Dear ImGui] library.

# Differences with other Deam ImGui bindings

This package is manually written with the goal of being idiomatic. It
focus on simplicity and reduces the number of dependencies to the
minimum. Do not expect it to be complete, as funcionalities will be
added when required by other packages. However, the idea is to keep a
simple code base that makes it easy to extend it.

# Pre-compiled imgui static library

This package bundles a pre-compiled cimgui static library to make it
easier to use.

Developers of this package can run the following command to build the
pre-compiled cimgui static library for the running host as well as the
corresponding header files:

	go generate

# Wayland

If you are targeting wayland, specify the "wayland" build tag. For
instance,

	go run -tags wayland github.com/jroimartin/imgui/examples/glfw-opengl@latest

[Dear ImGui]: https://github.com/ocornut/imgui
*/
package imgui
