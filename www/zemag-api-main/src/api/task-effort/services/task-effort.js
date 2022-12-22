'use strict';

/**
 * task-effort service.
 */

const { createCoreService } = require('@strapi/strapi').factories;

module.exports = createCoreService('api::task-effort.task-effort');
