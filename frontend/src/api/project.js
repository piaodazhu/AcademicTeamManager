import request from '../axios/index'

export function createProject(param) {
    return request({
		url: '/project/create',
		method: 'post',
		data: param,
	})
}

export function updateProject(param) {
    return request({
		url: '/project/update',
		method: 'post',
		data: param,
	})
}

export function deleteProject(param) {
    return request({
		url: '/project/delete',
		method: 'post',
		data: param,
	})
}


export function queryProjectList(param) {
    return request({
		url: '/project/list',
		method: 'get',
		params: param,
	})
}


export function queryProjectInfo(param) {
    return request({
		url: '/project/info',
		method: 'get',
		params: param,
	})
}

export function queryProjectOlist(param) {
    return request({
		url: '/project/olist',
		method: 'post',
		data: param,
	})
}


export function projectExport(param) {
    return request({
		url: '/project/export',
		method: 'get',
		responseType: 'blob',
		params: param,
	})
}