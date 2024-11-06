export namespace database {
	
	export class Brb {
	    ID: number;
	    CreatedAt: time.Time;
	    UpdatedAt: time.Time;
	    // Go type: gorm
	    DeletedAt: any;
	    ID: number;
	    TotalTimeID: number;
	    StartTime: time.Time;
	    Duration: number;
	    Active: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Brb(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], time.Time);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], time.Time);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.ID = source["ID"];
	        this.TotalTimeID = source["TotalTimeID"];
	        this.StartTime = this.convertValues(source["StartTime"], time.Time);
	        this.Duration = source["Duration"];
	        this.Active = source["Active"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class BreakTime {
	    ID: number;
	    CreatedAt: time.Time;
	    UpdatedAt: time.Time;
	    // Go type: gorm
	    DeletedAt: any;
	    ID: number;
	    TotalTimeID: number;
	    StartTime: time.Time;
	    Duration: number;
	    Active: boolean;
	
	    static createFrom(source: any = {}) {
	        return new BreakTime(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], time.Time);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], time.Time);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.ID = source["ID"];
	        this.TotalTimeID = source["TotalTimeID"];
	        this.StartTime = this.convertValues(source["StartTime"], time.Time);
	        this.Duration = source["Duration"];
	        this.Active = source["Active"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class Cost {
	    ID: number;
	    CreatedAt: time.Time;
	    UpdatedAt: time.Time;
	    // Go type: gorm
	    DeletedAt: any;
	    ID: number;
	    ProjectID: number;
	    Duration: number;
	    HourCost: number;
	
	    static createFrom(source: any = {}) {
	        return new Cost(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], time.Time);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], time.Time);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.ID = source["ID"];
	        this.ProjectID = source["ProjectID"];
	        this.Duration = source["Duration"];
	        this.HourCost = source["HourCost"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class Task {
	    ID: number;
	    CreatedAt: time.Time;
	    UpdatedAt: time.Time;
	    // Go type: gorm
	    DeletedAt: any;
	    ID: number;
	    ProjectID: number;
	    Deadline: time.Time;
	    Description: string;
	    Closed: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Task(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], time.Time);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], time.Time);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.ID = source["ID"];
	        this.ProjectID = source["ProjectID"];
	        this.Deadline = this.convertValues(source["Deadline"], time.Time);
	        this.Description = source["Description"];
	        this.Closed = source["Closed"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class WorkTime {
	    ID: number;
	    CreatedAt: time.Time;
	    UpdatedAt: time.Time;
	    // Go type: gorm
	    DeletedAt: any;
	    ID: number;
	    TotalTimeID: number;
	    StartTime: time.Time;
	    Duration: number;
	    Projects: Project[];
	    Closed: boolean;
	    Trustworthy: boolean;
	
	    static createFrom(source: any = {}) {
	        return new WorkTime(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], time.Time);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], time.Time);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.ID = source["ID"];
	        this.TotalTimeID = source["TotalTimeID"];
	        this.StartTime = this.convertValues(source["StartTime"], time.Time);
	        this.Duration = source["Duration"];
	        this.Projects = this.convertValues(source["Projects"], Project);
	        this.Closed = source["Closed"];
	        this.Trustworthy = source["Trustworthy"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class Project {
	    ID: number;
	    CreatedAt: time.Time;
	    UpdatedAt: time.Time;
	    // Go type: gorm
	    DeletedAt: any;
	    ID: number;
	    Name: string;
	    StartTime: time.Time;
	    Duration: number;
	    Closed: boolean;
	    Cost?: Cost;
	    WorkTimes: WorkTime[];
	    Tasks: Task[];
	
	    static createFrom(source: any = {}) {
	        return new Project(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], time.Time);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], time.Time);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.StartTime = this.convertValues(source["StartTime"], time.Time);
	        this.Duration = source["Duration"];
	        this.Closed = source["Closed"];
	        this.Cost = this.convertValues(source["Cost"], Cost);
	        this.WorkTimes = this.convertValues(source["WorkTimes"], WorkTime);
	        this.Tasks = this.convertValues(source["Tasks"], Task);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class ResolutionUnit {
	    ID: number;
	    CreatedAt: time.Time;
	    UpdatedAt: time.Time;
	    // Go type: gorm
	    DeletedAt: any;
	    ID: number;
	    TrackerID: number;
	    Tracker: ResolutionTracker;
	    Identifier: string;
	    Resolved: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ResolutionUnit(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], time.Time);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], time.Time);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.ID = source["ID"];
	        this.TrackerID = source["TrackerID"];
	        this.Tracker = this.convertValues(source["Tracker"], ResolutionTracker);
	        this.Identifier = source["Identifier"];
	        this.Resolved = source["Resolved"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class ResolutionTracker {
	    ID: number;
	    CreatedAt: time.Time;
	    UpdatedAt: time.Time;
	    // Go type: gorm
	    DeletedAt: any;
	    ID: number;
	    Day: time.Time;
	    Category: string;
	    Closed: boolean;
	    Units: ResolutionUnit[];
	
	    static createFrom(source: any = {}) {
	        return new ResolutionTracker(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], time.Time);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], time.Time);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.ID = source["ID"];
	        this.Day = this.convertValues(source["Day"], time.Time);
	        this.Category = source["Category"];
	        this.Closed = source["Closed"];
	        this.Units = this.convertValues(source["Units"], ResolutionUnit);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	
	export class SearchResult {
	    type: string;
	    id: number;
	    name?: string;
	    identifier?: string;
	    day?: time.Time;
	
	    static createFrom(source: any = {}) {
	        return new SearchResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.id = source["id"];
	        this.name = source["name"];
	        this.identifier = source["identifier"];
	        this.day = this.convertValues(source["day"], time.Time);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	
	export class TotalTime {
	    ID: number;
	    CreatedAt: time.Time;
	    UpdatedAt: time.Time;
	    // Go type: gorm
	    DeletedAt: any;
	    ID: number;
	    StartTime: time.Time;
	    FinishTime: time.Time;
	    WorkTimes: WorkTime[];
	    BreakTime?: BreakTime;
	    Brb?: Brb;
	    Closed: boolean;
	
	    static createFrom(source: any = {}) {
	        return new TotalTime(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], time.Time);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], time.Time);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.ID = source["ID"];
	        this.StartTime = this.convertValues(source["StartTime"], time.Time);
	        this.FinishTime = this.convertValues(source["FinishTime"], time.Time);
	        this.WorkTimes = this.convertValues(source["WorkTimes"], WorkTime);
	        this.BreakTime = this.convertValues(source["BreakTime"], BreakTime);
	        this.Brb = this.convertValues(source["Brb"], Brb);
	        this.Closed = source["Closed"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	
	export class WorkTimeProject {
	    ID: number;
	    WorkTimeID: number;
	    ProjectID: number;
	    StartTime: time.Time;
	    Duration: number;
	    Closed: boolean;
	    WorkTime: WorkTime;
	    Project: Project;
	
	    static createFrom(source: any = {}) {
	        return new WorkTimeProject(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.WorkTimeID = source["WorkTimeID"];
	        this.ProjectID = source["ProjectID"];
	        this.StartTime = this.convertValues(source["StartTime"], time.Time);
	        this.Duration = source["Duration"];
	        this.Closed = source["Closed"];
	        this.WorkTime = this.convertValues(source["WorkTime"], WorkTime);
	        this.Project = this.convertValues(source["Project"], Project);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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

export namespace main {
	
	export class CurrentTimersResponse {
	    workTime?: database.WorkTime;
	    breakTime?: database.BreakTime;
	    brb?: database.Brb;
	    totalTime?: database.TotalTime;
	
	    static createFrom(source: any = {}) {
	        return new CurrentTimersResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.workTime = this.convertValues(source["workTime"], database.WorkTime);
	        this.breakTime = this.convertValues(source["breakTime"], database.BreakTime);
	        this.brb = this.convertValues(source["brb"], database.Brb);
	        this.totalTime = this.convertValues(source["totalTime"], database.TotalTime);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class DaySummary {
	    workTimesStarted: database.WorkTime[];
	    workTimesCrossingDays: database.WorkTime[];
	    totalTime: number;
	    workTimeProjects: database.WorkTimeProject[];
	    breaks: database.BreakTime[];
	    brbs: database.Brb[];
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new DaySummary(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.workTimesStarted = this.convertValues(source["workTimesStarted"], database.WorkTime);
	        this.workTimesCrossingDays = this.convertValues(source["workTimesCrossingDays"], database.WorkTime);
	        this.totalTime = source["totalTime"];
	        this.workTimeProjects = this.convertValues(source["workTimeProjects"], database.WorkTimeProject);
	        this.breaks = this.convertValues(source["breaks"], database.BreakTime);
	        this.brbs = this.convertValues(source["brbs"], database.Brb);
	        this.message = source["message"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class MessageCostResponse {
	    message: string;
	    cost?: database.Cost;
	
	    static createFrom(source: any = {}) {
	        return new MessageCostResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	        this.cost = this.convertValues(source["cost"], database.Cost);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class MessageProjectResponse {
	    message: string;
	    project?: database.Project;
	
	    static createFrom(source: any = {}) {
	        return new MessageProjectResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	        this.project = this.convertValues(source["project"], database.Project);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class MessageWorkTimeProjectResponse {
	    message: string;
	    project?: database.Project;
	    workTimeProject?: database.WorkTimeProject;
	
	    static createFrom(source: any = {}) {
	        return new MessageWorkTimeProjectResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	        this.project = this.convertValues(source["project"], database.Project);
	        this.workTimeProject = this.convertValues(source["workTimeProject"], database.WorkTimeProject);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class MessageWorkTimeResponse {
	    message: string;
	    workTime?: database.WorkTime;
	
	    static createFrom(source: any = {}) {
	        return new MessageWorkTimeResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	        this.workTime = this.convertValues(source["workTime"], database.WorkTime);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class ProjectsResponse {
	    projects: database.Project[];
	    total: number;
	    currentPage: number;
	    itemsPerPage: number;
	
	    static createFrom(source: any = {}) {
	        return new ProjectsResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.projects = this.convertValues(source["projects"], database.Project);
	        this.total = source["total"];
	        this.currentPage = source["currentPage"];
	        this.itemsPerPage = source["itemsPerPage"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class ResolutionMessageResponse {
	    message: string;
	    tracker?: database.ResolutionTracker;
	    units?: database.ResolutionUnit[];
	    unit?: database.ResolutionUnit;
	    success: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ResolutionMessageResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	        this.tracker = this.convertValues(source["tracker"], database.ResolutionTracker);
	        this.units = this.convertValues(source["units"], database.ResolutionUnit);
	        this.unit = this.convertValues(source["unit"], database.ResolutionUnit);
	        this.success = source["success"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class SearchResultResponse {
	    success: boolean;
	    message: string;
	    results: database.SearchResult[];
	
	    static createFrom(source: any = {}) {
	        return new SearchResultResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.message = source["message"];
	        this.results = this.convertValues(source["results"], database.SearchResult);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class StartDayResponse {
	    message: string;
	    totalTime?: database.TotalTime;
	
	    static createFrom(source: any = {}) {
	        return new StartDayResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	        this.totalTime = this.convertValues(source["totalTime"], database.TotalTime);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class TaskResponse {
	    message: string;
	    task?: database.Task;
	
	    static createFrom(source: any = {}) {
	        return new TaskResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	        this.task = this.convertValues(source["task"], database.Task);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class Timer {
	    Duration: number;
	    Message: string;
	    FinishTime: time.Time;
	
	    static createFrom(source: any = {}) {
	        return new Timer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Duration = source["Duration"];
	        this.Message = source["Message"];
	        this.FinishTime = this.convertValues(source["FinishTime"], time.Time);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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

export namespace time {
	
	export class Time {
	
	
	    static createFrom(source: any = {}) {
	        return new Time(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

