import request from '../axios/index'

export function createStudent(param) {
    return request({
		url: '/student/create',
		method: 'post',
		data: param,
	})
}

export function updateStudent(param) {
    return request({
		url: '/student/update',
		method: 'post',
		data: param,
	})
}

export function sendMailToStudent(param) {
    return request({
		url: '/student/send',
		method: 'post',
		data: param,
	})
}

export function deleteStudent(param) {
    return request({
		url: '/student/delete',
		method: 'post',
		data: param,
	})
}

export function queryStudentList(param) {
    return request({
		url: '/student/list',
		method: 'get',
		params: param,
	})
}

export function queryStudentInfo(param) {
    return request({
		url: '/student/info',
		method: 'get',
		params: param,
	})
}


export function queryStudentOption(param) {
    return request({
		url: '/student/option',
		method: 'get',
		params: param,
	})
}

export function studentExport(param) {
    return request({
		url: '/student/export',
		method: 'get',
		responseType: 'blob',
		params: param,
	})
}