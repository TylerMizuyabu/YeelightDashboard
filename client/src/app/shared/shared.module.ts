import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SideMenuComponent } from './components/side-menu/side-menu.component';
import { MaterialModule } from '../material.module';



@NgModule({
  declarations: [SideMenuComponent],
  imports: [
    CommonModule,
    MaterialModule
  ],
  exports: [
    SideMenuComponent
  ]
})
export class SharedModule { }
