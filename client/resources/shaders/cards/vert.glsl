#version 330 core

layout(location = 0) in vec3 in_Position;
layout(location = 1) in vec2 in_TexCoords;

out vec2 v_TexCoords;

uniform mat4 u_Projection;
uniform mat4 u_Camera;
uniform mat4 u_Transform;

void main() {
  vec4 projected =
    u_Projection *
    u_Camera *
    u_Transform *
    vec4(in_Position, 1);

  gl_Position = projected;
  v_TexCoords = in_TexCoords;
};
