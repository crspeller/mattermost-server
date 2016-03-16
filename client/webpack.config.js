const webpack = require('webpack');
module.exports = {
    entry: ['babel-polyfill', './root.jsx'],
    output: {
        path: 'dist',
        publicPath: '/static/',
        filename: 'testwbundle.js'
    },
    devtool: 'inline-source-map',
    module: {
        loaders: [
            {
                test: /\.jsx?$/,
                loader: 'babel-loader',
                exclude: /(node_modules|non_npm_dependencies)/,
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
            },
            {
                test: /\.(png|eot|tiff|svg|woff2|woff|ttf)$/,
                loader: 'file-loader',
                query: {
                    name: 'files/[hash].[ext]'
                }
            }
        ]
    },
    sassLoader: {
        includePaths: ['node_modules/compass-mixins/lib']
    },
    plugins: [
        new webpack.ProvidePlugin({
            'window.jQuery': 'jquery'
        })
    ],
    resolve: {
        alias: {
            jquery: 'jquery/dist/jquery'
        },
        modules: [
            'node_modules',
            'non_npm_dependencies'
        ]
    }
};
