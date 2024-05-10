#version 330 core

layout(location = 0) in vec3 position;
layout(location = 1) in vec2 texCoords;

out vec2 v_TexCoords;

uniform mat4 u_Transform;

void main() {
  vec4 transformed = vec4(position.xyz, 1) * u_Transform;

  gl_Position = vec4(transformed.xyz, 1);
  v_TexCoords = texCoords;
};
