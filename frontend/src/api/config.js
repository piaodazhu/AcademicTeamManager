import request from '../axios/index'

export function saveMailConfig(param) {
    return request({
		url: '/mailconf/save',
		method: 'post',
		data: param,
	})
}

export function deleteMailConfig(param) {
    return request({
		url: '/mailconf/delete',
		method: 'post',
		data: param,
	})
}

export function updateMailConfigStmpStatus(param) {
    return request({
		url: '/mailconf/switch/set',
		method: 'post',
		data: param,
	})
}

export function getMailConfig(param) {
    return request({
		url: '/mailconf/info',
		method: 'get',
		data: param,
	})
}

export function checkMailConfig(param) {
    return request({
		url: '/mailconf/check',
		method: 'get',
		data: param,
	})
}