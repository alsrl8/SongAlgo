// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {github} from '../models';
import {selenium} from '../models';

export function GetGithubRepositoryBjSource(arg1:string,arg2:string,arg3:string,arg4:string):Promise<github.FileResponse>;

export function GetMenu():Promise<any>;

export function GetSchedule():Promise<github.ScheduleList>;

export function IsChromeRunning():Promise<boolean>;

export function NavigateToBjProblemWithCookie(arg1:string):Promise<Array<selenium.SubmitHistory>>;

export function UploadBjSourceToGithub(arg1:string,arg2:string,arg3:selenium.SubmitHistory,arg4:string):Promise<void>;
