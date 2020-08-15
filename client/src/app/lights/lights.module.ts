import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LightListComponent } from './components/light-list/light-list.component';
import {MaterialModule} from '../material.module';



@NgModule({
  declarations: [LightListComponent],
  imports: [
    CommonModule,
    MaterialModule
  ]
})
export class LightsModule { }
