module.exports = ({ env }) => ({
  auth: {
    secret: env('ADMIN_JWT_SECRET', ''),
    secret: env('API_TOKEN_SALT', '')
  }
});
