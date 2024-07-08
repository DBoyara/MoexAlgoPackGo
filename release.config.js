// release.config.js
module.exports = {
    branches: ['master'],
    plugins: [
      '@semantic-release/commit-analyzer',
      '@semantic-release/release-notes-generator',
      '@semantic-release/changelog',
      ['@semantic-release/git', {
        assets: ['CHANGELOG.md', 'go.mod', 'go.sum'],
        message: 'chore(release): ${nextRelease.version} [skip ci]\n\n${nextRelease.notes}'
      }],
      '@semantic-release/github'
    ],
  };