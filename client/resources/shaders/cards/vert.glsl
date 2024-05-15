#version 330 core

layout(location = 0) in vec3 in_Position;
layout(location = 1) in vec2 in_TexCoords;

out vec3 v_TexCoords;
out float v_Depth;

uniform mat4 u_Projection;
uniform mat4 u_TransformScale;
uniform mat4 u_TransformRotation;
uniform mat4 u_TransformTranslation;

void main() {
  mat4 transform = u_TransformTranslation * u_TransformRotation * u_TransformScale;
  vec4 projected = u_Projection * transform * vec4(in_Position.xyz, 1);
  float depth = projected.w;
  vec4 perspectived = projected / depth;

  gl_Position = vec4(perspectived.xyz, 1);
  v_Depth = depth;

  // Dividing to ensure proper interpolation.
  // Should be multiplied back together in the fragment shader.
  v_TexCoords = vec3(in_TexCoords / depth, 1/depth);
};
