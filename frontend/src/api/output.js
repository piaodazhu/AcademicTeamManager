import request from '../axios/index'

export function createOutput(param) {
    return request({
		url: '/output/create',
		method: 'post',
		data: param,
	})
}

export function updateOutput(param) {
    return request({
		url: '/output/update',
		method: 'post',
		data: param,
	})
}

export function deleteOutput(param) {
    return request({
		url: '/output/delete',
		method: 'post',
		data: param,
	})
}

export function queryOutputInfo(param) {
    return request({
		url: '/output/info',
		method: 'get',
		params: param,
	})
}

export function queryOutputList(param) {
    return request({
		url: '/output/list',
		method: 'get',
		params: param,
	})
}

export function outputExport(param) {
    return request({
		url: '/output/export',
		method: 'get',
		responseType: 'blob',
		params: param,
	})
}