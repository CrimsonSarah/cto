#version 330 core

uniform sampler2D sprite;

in vec3 v_TexCoords;
in float v_Depth;

layout(location = 0) out vec4 color;

void main() {
  color = textureProj(sprite, v_TexCoords);
};
