import { Spectator, createComponentFactory } from '@ngneat/spectator';

import { SideMenuComponent } from './side-menu.component';

describe('SideMenuComponent', () => {
  let spectator: Spectator<SideMenuComponent>;
  const createComponent = createComponentFactory(SideMenuComponent);

  it('should create', () => {
    spectator = createComponent();

    expect(spectator.component).toBeTruthy();
  });
});
