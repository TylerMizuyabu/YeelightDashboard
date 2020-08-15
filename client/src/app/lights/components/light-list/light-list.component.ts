import { Component, OnInit } from '@angular/core';
import { Store, select } from '@ngrx/store';
import * as lightSelectors from '../../../shared/selectors/lights.selector'
import { Observable } from 'rxjs';
import { Yeelight } from 'src/app/shared/models/light';

@Component({
  selector: 'app-light-list',
  templateUrl: './light-list.component.html',
  styleUrls: ['./light-list.component.scss']
})
export class LightListComponent implements OnInit {
  lights: Observable<Yeelight[]>;
  constructor(private store: Store) { }

  ngOnInit(): void {
    this.lights = this.store.pipe(select(lightSelectors.selectAllLights))
  }

}
