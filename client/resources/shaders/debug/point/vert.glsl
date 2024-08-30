#version 330 core

layout(location = 0) in vec3 in_Position;

out vec3 v_Color;
out vec3 v_Pos; // Position in the same coordinate system as u_Center.

uniform mat4 u_Projection;
uniform mat4 u_Transform;
uniform vec3 u_Color;

void main() {
  vec4 projected =
      u_Projection *
      u_Transform *
      vec4(in_Position, 1);

  gl_Position = projected;
  v_Color = u_Color;
  v_Pos = in_Position;
};
