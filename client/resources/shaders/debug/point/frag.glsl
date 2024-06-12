#version 330 core

in vec3 v_Color;

uniform vec2 u_Center;
uniform float u_Radius;
uniform vec2 u_WindowDimensions;

layout(location = 0) out vec4 color;

void main() {
  // Components between 0 and 1.
  vec2 coords = gl_FragCoord.xy / u_WindowDimensions;
  // Components between -1 and +1.
  coords = coords * 2 - vec2(1, 1);

  vec2 diff = coords - u_Center.xy;
  float dist = diff.x * diff.x + diff.y * diff.y;

  if (dist * dist < u_Radius * u_Radius) {
    color = vec4(v_Color, 1);
  } else {
    color = vec4(1, 0, 0, 0);
  }
};
