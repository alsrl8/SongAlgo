export namespace github {
	
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
	
	export class SubmitHistory {
	    "제출 번호": string;
	    "아이디": string;
	    "문제": string;
	    "결과": string;
	    "메모리": string;
	    "시간": string;
	    "언어": string;
	    "코드 길이": string;
	    "제출한 시간": string;
	
	    static createFrom(source: any = {}) {
	        return new SubmitHistory(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this["제출 번호"] = source["제출 번호"];
	        this["아이디"] = source["아이디"];
	        this["문제"] = source["문제"];
	        this["결과"] = source["결과"];
	        this["메모리"] = source["메모리"];
	        this["시간"] = source["시간"];
	        this["언어"] = source["언어"];
	        this["코드 길이"] = source["코드 길이"];
	        this["제출한 시간"] = source["제출한 시간"];
	    }
	}

}

