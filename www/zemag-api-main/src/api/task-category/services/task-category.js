'use strict';

/**
 * task-category service.
 */

const { createCoreService } = require('@strapi/strapi').factories;

module.exports = createCoreService('api::task-category.task-category');
