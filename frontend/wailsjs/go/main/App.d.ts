// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {main} from '../models';
import {time} from '../models';
import {database} from '../models';

export function AssociateProjectToWorkTime(arg1:number):Promise<main.MessageWorkTimeProjectResponse>;

export function CalculateAndSaveProjectCost(arg1:number,arg2:number):Promise<main.MessageCostResponse>;

export function CalculateWorkTimeForDay(arg1:time.Time):Promise<time.Duration>;

export function ChangeProjectClose(arg1:number):Promise<main.MessageProjectResponse>;

export function CreateProject(arg1:string):Promise<main.MessageProjectResponse>;

export function CreateTask(arg1:number,arg2:string,arg3:string):Promise<main.TaskResponse>;

export function EndBreak():Promise<main.MessageWorkTimeResponse>;

export function FinishDay():Promise<string>;

export function GetAllProjects(arg1:number,arg2:number,arg3:any,arg4:string,arg5:string):Promise<main.ProjectsResponse>;

export function GetDaySummary(arg1:time.Time):Promise<main.DaySummary>;

export function GetProjectByID(arg1:number):Promise<main.MessageProjectResponse>;

export function GetStartTimes():Promise<Array<main.Timer>>;

export function GetUnfinishedWorkTimeProjectWithoutSendingError():Promise<database.WorkTimeProject>;

export function RemoveTimer(arg1:string):Promise<string>;

export function StartDay():Promise<main.StartDayResponse>;

export function StartTimer(arg1:number,arg2:string):Promise<string>;

export function StartWorkTime():Promise<main.MessageWorkTimeResponse>;

export function TakeBreak():Promise<string>;

export function UpdateProjectName(arg1:number,arg2:string):Promise<main.MessageProjectResponse>;
