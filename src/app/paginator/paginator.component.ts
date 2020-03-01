import { ChangeDetectionStrategy, Component, EventEmitter, Input, OnChanges, Output } from '@angular/core';

import { BehaviorSubject } from 'rxjs';

@Component({
  selector: 'survey-paginator',
  templateUrl: './paginator.component.html',
  styleUrls: ['./paginator.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PaginatorComponent implements OnChanges {
  @Input() public maxPages: number;
  @Input() public currentPage: number;
  @Input() public disabledNextBtn: boolean;

  @Output() public onPage: EventEmitter<number> = new EventEmitter();

  public progress$: BehaviorSubject<number> = new BehaviorSubject(0);

  ngOnChanges() {
    const percentage: number = Math.ceil(((this.currentPage + 1) * 100) / this.maxPages);
    this.progress$.next(percentage);
  }

  public handlePage(page: number): void {
    if ((page > this.currentPage && !this.disabledNextBtn) || (page >= 0 && page < this.maxPages)) {
      this.onPage.emit(page);
    }
  }
}
