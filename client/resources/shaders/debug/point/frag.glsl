#version 330 core

in vec3 v_Color;
in vec3 v_Pos;

uniform vec3 u_Center;
uniform float u_Radius;
uniform vec2 u_WindowDimensions;

layout(location = 0) out vec4 color;

void main() {
  vec3 diff = v_Pos - u_Center;
  float dist = diff.x * diff.x + diff.y * diff.y;

  if (dist < u_Radius * u_Radius) {
    color = vec4(v_Color, 1);
  } else {
    color = vec4(1, 0, 0, 0);
  }
};
