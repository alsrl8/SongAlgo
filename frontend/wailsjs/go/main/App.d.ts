// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {github} from '../models';
import {selenium} from '../models';

export function AddProblem(arg1:string,arg2:string,arg3:string,arg4:string,arg5:string):Promise<void>;

export function CloseProgram():Promise<void>;

export function CloseSeleniumBrowser():Promise<void>;

export function GetGithubRepositoryBjSource(arg1:string,arg2:string,arg3:string,arg4:string):Promise<github.FileResponse>;

export function GetGithubRepositoryPgSource(arg1:string,arg2:string,arg3:string,arg4:string):Promise<github.FileResponse>;

export function GetMenu():Promise<any>;

export function GetPgSourceData(arg1:string):Promise<selenium.PgSourceData>;

export function GetSchedule():Promise<github.ScheduleList>;

export function IsBjLoggedIn(arg1:string):Promise<boolean>;

export function IsChromeRunning():Promise<boolean>;

export function IsPgLoggedIn(arg1:string):Promise<boolean>;

export function IsSubmittedCodeCorrect(arg1:string):Promise<boolean>;

export function NavigateToBjLoginPage():Promise<void>;

export function NavigateToBjProblemWithCookie(arg1:string):Promise<Array<selenium.SubmitHistory>>;

export function NavigateToPgLoginPage():Promise<void>;

export function UploadBjSourceToGithub(arg1:string,arg2:string,arg3:selenium.SubmitHistory,arg4:string):Promise<void>;

export function UploadPgSourceToGithub(arg1:string,arg2:string,arg3:string,arg4:string,arg5:string,arg6:string):Promise<void>;
