import request from '../axios/index'

export function initDatabase(param) {
    return request({
		url: '/util/initdb',
		method: 'get',
		data: param,
	})
}

export function fileRemove(param) {
    return request({
		url: '/util/removefile',
		method: 'post',
		data: param,
	})
}

export function fileUpload(param) {
	return request({
		    url: '/util/uploadfile',
		    method: 'post',
		    data: param,
	    })
    }
