export * from './mipt.service';
import { MiptService } from './mipt.service';
export * from './spreadsheet.service';
import { SpreadsheetService } from './spreadsheet.service';
export * from './state.service';
import { StateService } from './state.service';
export * from './users.service';
import { UsersService } from './users.service';
export const APIS = [MiptService, SpreadsheetService, StateService, UsersService];
