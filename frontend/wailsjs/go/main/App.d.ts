// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {database} from '../models';
import {main} from '../models';

export function EndBreak():Promise<string|database.WorkTime>;

export function FindByID(arg1:number,arg2:string):Promise<any|string>;

export function StartDay():Promise<main.StartDayResponse>;

export function StartTimer(arg1:number,arg2:string):Promise<string>;

export function StartWorkTime():Promise<main.StartWorkTimeResponse>;

export function TakeBreak(arg1:number):Promise<string>;
