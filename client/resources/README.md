# Resources

This is at the same time a Go package and a directory for resource files, like textures and shaders.

Go files live in the root and actual resource files live in the directories.

Resource objects should only be called after OpenGL has been initialized, since they do interact with it.
