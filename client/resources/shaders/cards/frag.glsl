#version 330 core

uniform sampler2D sprite;

in vec2 v_TexCoords;

layout(location = 0) out vec4 color;

void main() {
  color = texture(sprite, v_TexCoords);
};
