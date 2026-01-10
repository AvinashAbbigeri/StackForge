#!/usr/bin/env bash
set -e

sudo dnf install -y nodejs
npm install express
npm install -D eslint
npm install -D prettier
npm install -D jest
npm install -D typescript

mkdir -p $(dirname package.json)
cat << 'EOF' > package.json
{
  "name": "stackforge-app",
  "version": "1.0.0",
  "main": "index.js",
  "scripts": {
    "start": "node index.js"
  }
}

EOF
mkdir -p $(dirname index.js)
cat << 'EOF' > index.js
console.log('Hello from StackForge Node');
EOF
mkdir -p $(dirname index.js)
cat << 'EOF' > index.js
const express = require('express');
const app = express();

app.get('/', (req, res) => {
  res.send('Hello from StackForge Express');
});

app.listen(3000, () => console.log('Server running on port 3000'));

EOF
mkdir -p $(dirname eslint.config.js)
cat << 'EOF' > eslint.config.js
export default [
  {
    files: ['**/*.js'],
    rules: {
      'no-unused-vars': 'warn',
      'no-undef': 'error'
    }
  }
]

EOF
mkdir -p $(dirname .prettierrc)
cat << 'EOF' > .prettierrc
{
  "semi": false,
  "singleQuote": true
}

EOF
mkdir -p $(dirname test/basic.test.js)
cat << 'EOF' > test/basic.test.js
test('adds 1 + 1', () => { expect(1 + 1).toBe(2); });
EOF
mkdir -p $(dirname tsconfig.json)
cat << 'EOF' > tsconfig.json
{
  "compilerOptions": {
    "target": "ES2020",
    "module": "commonjs",
    "strict": true
  }
}

EOF

node --version
npm --version
node index.js
node index.js
npx eslint --version
npx prettier --version
npx jest --version
npx tsc --version

echo "StackForge setup complete"
