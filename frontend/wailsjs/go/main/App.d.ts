// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {main} from '../models';

export function AssociateProjectToWorkTime(arg1:number):Promise<main.MessageProjectResponse>;

export function CreateProject(arg1:string):Promise<main.MessageProjectResponse>;

export function EndBreak():Promise<main.MessageWorkTimeResponse>;

export function FinishDay():Promise<string>;

export function GetAllProjects(arg1:number,arg2:number):Promise<main.ProjectsResponse>;

export function GetProjectByID(arg1:number):Promise<main.MessageProjectResponse>;

export function GetStartTimes():Promise<Array<main.Timer>>;

export function RemoveTimer(arg1:string):Promise<void>;

export function StartDay():Promise<main.StartDayResponse>;

export function StartTimer(arg1:number,arg2:string):Promise<string>;

export function StartWorkTime():Promise<main.MessageWorkTimeResponse>;

export function TakeBreak():Promise<string>;
