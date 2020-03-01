import { Injectable } from '@angular/core';

import { Observable, of } from 'rxjs';

import { ConfigInterface } from './interfaces/config.interface';
import { mockConfig } from './config.mock';
import {MiptService} from "../../swagger";
import {map} from "rxjs/operators";

@Injectable({
  providedIn: 'root',
})
export class ConfigService {
  private mockConfig: ConfigInterface = mockConfig;

  constructor(private miptService: MiptService) {
  }

  public getConfig(id: string) {
    return this.miptService.miptIdGet(id).pipe(
      map((r: any) => r.data)
    );
  }
}
