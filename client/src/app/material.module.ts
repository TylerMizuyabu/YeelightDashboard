

import {MatSidenavModule} from '@angular/material/sidenav'; 
import { NgModule } from '@angular/core';

const modules = [
  MatSidenavModule
]

@NgModule ({
  imports: [
    ...modules
  ],
  exports: [
    ...modules
  ]
})
export class MaterialModule{}