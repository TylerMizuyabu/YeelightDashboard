import { Component, OnInit, Input } from '@angular/core';
import {NestedTreeControl} from '@angular/cdk/tree';
import {MatTreeNestedDataSource} from '@angular/material/tree';
import {MenuItem} from './menu-item';

@Component({
  selector: 'app-side-menu',
  templateUrl: './side-menu.component.html',
  styleUrls: ['./side-menu.component.scss']
})
export class SideMenuComponent implements OnInit {
  treeControl = new NestedTreeControl<MenuItem>(item => item.children);
  datasource = new MatTreeNestedDataSource<MenuItem>();

  constructor() { }

  ngOnInit(): void {
  }

  @Input() set menuItems(items: MenuItem[]) {
    this.datasource.data = [...items]
  }

  hasChild(_: number, node: MenuItem) {
    return !!node.children && node.children.length > 0;
  }
}
