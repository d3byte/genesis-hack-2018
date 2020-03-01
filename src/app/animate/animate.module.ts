import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';

import { AnimateComponentTrigger } from './animate.component';

@NgModule({
  declarations: [AnimateComponentTrigger],
  imports: [CommonModule],
  exports: [AnimateComponentTrigger],
})
export class AnimateModule {}
