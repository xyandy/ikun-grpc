export namespace backend {
	
	export class CommonResp {
	
	
	    static createFrom(source: any = {}) {
	        return new CommonResp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

