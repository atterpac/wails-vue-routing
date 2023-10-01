export namespace models {
	
	export class Accessory {
	    id: string;
	    device_type: string;
	    symbol_name: string;
	    instrument_type: string;
	    fixture_id: number;
	
	    static createFrom(source: any = {}) {
	        return new Accessory(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.device_type = source["device_type"];
	        this.symbol_name = source["symbol_name"];
	        this.instrument_type = source["instrument_type"];
	        this.fixture_id = source["fixture_id"];
	    }
	}
	export class Fixture {
	    id: number[];
	    is_vwx?: boolean;
	    uid?: string;
	    is_gdtf?: boolean;
	    gdtf_uid?: string;
	    instrument_type?: string;
	    wattage?: string;
	    purpose?: string;
	    position?: string;
	    unit_number?: string;
	    color?: string;
	    dimmer?: string;
	    channel?: string;
	    universe?: string;
	    u_dimmer?: string;
	    circuit_number?: string;
	    circuit_name?: string;
	    system?: string;
	    num_channels?: string;
	    weight?: string;
	    focus?: string;
	    fixture_id?: string;
	    address?: string;
	    voltage?: string;
	    fixture_mode?: string;
	    accessories: Accessory[];
	    x_loc_mm?: number;
	    y_loc_mm?: number;
	    z_loc_mm?: number;
	    event_id: number[];
	
	    static createFrom(source: any = {}) {
	        return new Fixture(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.is_vwx = source["is_vwx"];
	        this.uid = source["uid"];
	        this.is_gdtf = source["is_gdtf"];
	        this.gdtf_uid = source["gdtf_uid"];
	        this.instrument_type = source["instrument_type"];
	        this.wattage = source["wattage"];
	        this.purpose = source["purpose"];
	        this.position = source["position"];
	        this.unit_number = source["unit_number"];
	        this.color = source["color"];
	        this.dimmer = source["dimmer"];
	        this.channel = source["channel"];
	        this.universe = source["universe"];
	        this.u_dimmer = source["u_dimmer"];
	        this.circuit_number = source["circuit_number"];
	        this.circuit_name = source["circuit_name"];
	        this.system = source["system"];
	        this.num_channels = source["num_channels"];
	        this.weight = source["weight"];
	        this.focus = source["focus"];
	        this.fixture_id = source["fixture_id"];
	        this.address = source["address"];
	        this.voltage = source["voltage"];
	        this.fixture_mode = source["fixture_mode"];
	        this.accessories = this.convertValues(source["accessories"], Accessory);
	        this.x_loc_mm = source["x_loc_mm"];
	        this.y_loc_mm = source["y_loc_mm"];
	        this.z_loc_mm = source["z_loc_mm"];
	        this.event_id = source["event_id"];
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
	export class FollowspotCue {
	    id: number[];
	    number: number;
	    frame: number;
	    size: string;
	    target: string;
	    description: string;
	    location: string;
	    notes: string;
	    followspot_id: number[];
	    followspot: Followspot;
	
	    static createFrom(source: any = {}) {
	        return new FollowspotCue(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.number = source["number"];
	        this.frame = source["frame"];
	        this.size = source["size"];
	        this.target = source["target"];
	        this.description = source["description"];
	        this.location = source["location"];
	        this.notes = source["notes"];
	        this.followspot_id = source["followspot_id"];
	        this.followspot = this.convertValues(source["followspot"], Followspot);
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
	export class Frame {
	    id: number[];
	    index: number;
	    label: string;
	    number: number;
	    followspot_id: number[];
	
	    static createFrom(source: any = {}) {
	        return new Frame(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.index = source["index"];
	        this.label = source["label"];
	        this.number = source["number"];
	        this.followspot_id = source["followspot_id"];
	    }
	}
	export class Followspot {
	    id: number[];
	    number: number;
	    fixture: string;
	    location: string;
	    notes: string;
	    frames: Frame[];
	    event_id: number[];
	    followspot_cues: FollowspotCue[];
	    // Go type: time
	    created_at: any;
	    // Go type: time
	    modified_at: any;
	    // Go type: gorm
	    deleted_at: any;
	
	    static createFrom(source: any = {}) {
	        return new Followspot(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.number = source["number"];
	        this.fixture = source["fixture"];
	        this.location = source["location"];
	        this.notes = source["notes"];
	        this.frames = this.convertValues(source["frames"], Frame);
	        this.event_id = source["event_id"];
	        this.followspot_cues = this.convertValues(source["followspot_cues"], FollowspotCue);
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.modified_at = this.convertValues(source["modified_at"], null);
	        this.deleted_at = this.convertValues(source["deleted_at"], null);
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
	export class User {
	    id: number[];
	    first_name: string;
	    last_name: string;
	    username: string;
	    email: string;
	    password: string;
	    events: Event[];
	    // Go type: time
	    modified_at: any;
	    // Go type: time
	    created_at: any;
	    // Go type: gorm
	    deleted_at: any;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.first_name = source["first_name"];
	        this.last_name = source["last_name"];
	        this.username = source["username"];
	        this.email = source["email"];
	        this.password = source["password"];
	        this.events = this.convertValues(source["events"], Event);
	        this.modified_at = this.convertValues(source["modified_at"], null);
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.deleted_at = this.convertValues(source["deleted_at"], null);
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
	export class Event {
	    id: number[];
	    name: string;
	    location: string;
	    start_date: string;
	    end_date: string;
	    users: User[];
	    followspots: Followspot[];
	    fixtures: Fixture[];
	    // Go type: gorm
	    deleted_at: any;
	    // Go type: time
	    created_at: any;
	    // Go type: time
	    modified_at: any;
	
	    static createFrom(source: any = {}) {
	        return new Event(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.location = source["location"];
	        this.start_date = source["start_date"];
	        this.end_date = source["end_date"];
	        this.users = this.convertValues(source["users"], User);
	        this.followspots = this.convertValues(source["followspots"], Followspot);
	        this.fixtures = this.convertValues(source["fixtures"], Fixture);
	        this.deleted_at = this.convertValues(source["deleted_at"], null);
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.modified_at = this.convertValues(source["modified_at"], null);
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

