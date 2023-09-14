export namespace github {
	
	export class Links {
	    self: string;
	    git: string;
	    html: string;
	
	    static createFrom(source: any = {}) {
	        return new Links(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.self = source["self"];
	        this.git = source["git"];
	        this.html = source["html"];
	    }
	}
	export class FileResponse {
	    name: string;
	    path: string;
	    sha: string;
	    size: number;
	    url: string;
	    html_url: string;
	    git_url: string;
	    download_url: string;
	    type: string;
	    content: string;
	    encoding: string;
	    // Go type: Links
	    _links: any;
	    statusCode: string;
	
	    static createFrom(source: any = {}) {
	        return new FileResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.path = source["path"];
	        this.sha = source["sha"];
	        this.size = source["size"];
	        this.url = source["url"];
	        this.html_url = source["html_url"];
	        this.git_url = source["git_url"];
	        this.download_url = source["download_url"];
	        this.type = source["type"];
	        this.content = source["content"];
	        this.encoding = source["encoding"];
	        this._links = this.convertValues(source["_links"], null);
	        this.statusCode = source["statusCode"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Problem {
	    name: string;
	    algorithmType: string;
	    difficulty: string;
	    platform: string;
	    url: string;
	
	    static createFrom(source: any = {}) {
	        return new Problem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.algorithmType = source["algorithmType"];
	        this.difficulty = source["difficulty"];
	        this.platform = source["platform"];
	        this.url = source["url"];
	    }
	}
	export class Schedule {
	    date: string;
	    problems: Problem[];
	
	    static createFrom(source: any = {}) {
	        return new Schedule(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.date = source["date"];
	        this.problems = this.convertValues(source["problems"], Problem);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ScheduleList {
	    list: Schedule[];
	
	    static createFrom(source: any = {}) {
	        return new ScheduleList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.list = this.convertValues(source["list"], Schedule);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace selenium {
	
	export class PgSourceData {
	    code: string;
	    extension: string;
	
	    static createFrom(source: any = {}) {
	        return new PgSourceData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.extension = source["extension"];
	    }
	}
	export class SubmitHistory {
	    SubmissionNumber: string;
	    ID: string;
	    Problem: string;
	    Result: string;
	    Memory: string;
	    Time: string;
	    Language: string;
	    CodeLength: string;
	    SubmissionTime: string;
	
	    static createFrom(source: any = {}) {
	        return new SubmitHistory(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.SubmissionNumber = source["SubmissionNumber"];
	        this.ID = source["ID"];
	        this.Problem = source["Problem"];
	        this.Result = source["Result"];
	        this.Memory = source["Memory"];
	        this.Time = source["Time"];
	        this.Language = source["Language"];
	        this.CodeLength = source["CodeLength"];
	        this.SubmissionTime = source["SubmissionTime"];
	    }
	}

}

