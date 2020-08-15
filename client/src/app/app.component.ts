import { Component } from '@angular/core';
import { MenuItem } from './shared/components/side-menu/menu-item';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'client';

  menuItems: MenuItem[] = [{
    icon: 'devices',
    text: 'Lights',
  }]
}
