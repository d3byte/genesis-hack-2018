import { Injectable } from '@angular/core';

import {MiptService} from "../../swagger";
import {map} from "rxjs/operators";

@Injectable({
  providedIn: 'root',
})
export class ConfigService {
  constructor(private miptService: MiptService) {
  }

  public getConfig(id: string) {
    return this.miptService.miptIdGet(id).pipe(
      map((r: any) => r.data)
    );
  }
}
