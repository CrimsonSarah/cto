#version 330 core

uniform sampler2D sprite;

in vec2 v_TexCoords;
in float v_Depth;

layout(location = 0) out vec4 color;

void main() {
  color = texture(sprite, v_TexCoords);
};
