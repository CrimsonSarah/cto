#version 330 core

layout(location = 0) in vec3 in_Position;
layout(location = 1) in vec2 in_TexCoords;

out vec2 v_TexCoords;
out float v_Depth;

uniform mat4 u_Projection;
uniform mat4 u_Transform;

void main() {
  vec4 transformed = vec4(in_Position.xyz, 1) * u_Transform * u_Projection;

  gl_Position = vec4(transformed.xyz, 1);

  v_TexCoords = in_TexCoords;
  v_Depth = transformed.w;
};
