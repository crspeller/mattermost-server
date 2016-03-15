module.exports = {
    entry: ['babel-polyfill', './pages/root.jsx'],
    output: {
        path: __dirname,
        filename: 'testwbundle.js'
    },
    devtool: 'inline-source-map',
    module: {
        loaders: [
            {
                test: /\.jsx?$/,
                loader: 'babel-loader',
                exclude: /node_modules/,
                query: {
                    presets: ['react', 'es2015', 'stage-0'],
                    plugins: ['transform-runtime']
                }
            },
            {
                test: /\.json$/,
                loader: 'json-loader'
            },
            {
                test: /(node_modules|non_npm_dependencies)\/.+\.(js|jsx)$/,
                loader: 'imports-loader',
                query: {
                    $: 'jquery',
                    jQuery: 'jquery'
                }
            },
            {
                test: /\.scss$/,
                loaders: ['style', 'css', 'sass']
            },
            {
                test: /\.css$/,
                loaders: ['style', 'css']
            }
        ]
    },
    sassLoader: {
        includePaths: ["node_modules/compass-mixins/lib"]
    },
    resolve: {
        alias: {
            jquery: 'jquery/dist/jquery'
        },
        modules: [
            "node_modules",
            "non_npm_dependencies"
        ]
    }
};
