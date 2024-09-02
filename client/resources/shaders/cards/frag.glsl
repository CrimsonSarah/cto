#version 330 core

uniform sampler2D texSprite;
uniform sampler2D texBack;

in vec2 v_TexCoords;

layout(location = 0) out vec4 color;

void main() {
  if (gl_FrontFacing) {
    color = texture(texSprite, v_TexCoords);
  } else {
    color = texture(texBack, vec2(-v_TexCoords.x, v_TexCoords.y));
  }
};
