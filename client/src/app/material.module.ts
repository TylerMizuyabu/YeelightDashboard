import { MatTreeModule } from '@angular/material/tree';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatIconModule } from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {MatTableModule} from '@angular/material/table';
import { NgModule } from '@angular/core';

const modules = [
  MatSidenavModule,
  MatTreeModule,
  MatIconModule,
  MatButtonModule,
  MatTableModule
]

@NgModule({
  imports: [
    ...modules
  ],
  exports: [
    ...modules
  ]
})
export class MaterialModule { }