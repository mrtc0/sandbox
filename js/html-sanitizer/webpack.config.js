module.exports = {
  mode: 'development',
  entry: './src/index.ts',
  target: 'web',
  module: {
    rules: [
      {
        test: /\.ts$/,
        use: 'ts-loader',
      },
    ],
  },
  resolve: {
    extensions: [
      '.ts', '.js',
    ],
  },
  output: {
    filename: 'dist.js',
    path: __dirname + '/dist',
    library: 'HtmlSanitizer',
    libraryTarget: 'window'
  }
};
