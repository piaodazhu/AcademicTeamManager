import request from '../axios/index'

export function updateNotice(param) {
    return request({
		url: '/notice/update',
		method: 'put',
		data: param,
	})
}

export function getNoticeCount(param) {
    return request({
		url: '/notice/count',
		method: 'get',
		params: param,
	})
}

export function deleteNotice(param) {
    return request({
		url: '/notice/delete',
		method: 'delete',
		data: param,
	})
}

export function getNoticeList(param) {
    return request({
		url: '/notice/list',
		method: 'get',
		params: param,
	})
}