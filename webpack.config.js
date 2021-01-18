const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
// 引入css 单独打包插件
const MiniCssExtractPlugin = require('mini-css-extract-plugin');

// 获取绝对路径
const resolve = dir=>path.resolve(__dirname, dir);

// 避免CSS重复复制
function recursiveIssuer(m, c) {
    const issuer = c.moduleGraph.getIssuer(m);
    // For webpack@4 chunks = m.issuer

    if (issuer) {
        return recursiveIssuer(issuer, c);
    }

    const chunks = c.chunkGraph.getModuleChunks(m);
    // For webpack@4 chunks = m._chunks

    for (const chunk of chunks) {
        return chunk.name;
    }

    return false;
}

module.exports = {
  mode:"development",
  entry: {
    index: path.resolve(__dirname, 'src/pages/index/index.js'),
    about: path.resolve(__dirname, 'src/pages/about/about.js')
  },
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'js/[name].js'
  },
  // 调试用，出错时直接定位到原始代码，而不是转换后的代码
  devtool:'cheap-module-eval-source-map',
    //   optimization: {
//     splitChunks: {
//       cacheGroups: {
//         indexStyles: {
//           name: 'styles_index',
//           test: (m, c, entry = 'index') =>
//             m.constructor.name === 'CssModule' &&
//             recursiveIssuer(m, c) === entry,
//           chunks: 'all',
//           enforce: true,
//         },
//         aboutStyles: {
//           name: 'styles_about',
//           test: (m, c, entry = 'about') =>
//             m.constructor.name === 'CssModule' &&
//             recursiveIssuer(m, c) === entry,
//           chunks: 'all',
//           enforce: true,
//         },
//       },
//     },
//   },
  resolve:{
    // 自动补全（可以省略）的扩展名
    extensions:['.js'],
    // 路径别名
    alias:{
        api: path.resolve(__dirname,'src/api'),
        fonts: path.resolve(__dirname,'src/assets/fonts'),
        images: path.resolve(__dirname,'src/assets/images'),
        styles: path.resolve(__dirname,'src/assets/styles'),
        components: path.resolve(__dirname,'src/components'),
        pages: path.resolve(__dirname,'src/pages'),
    }
},
  module: {
    rules: [{
        test: /\.(png|jpg|gif|svg)$/i,
        use:{        
            loader: "url-loader",
            options:{
                // 小于10K的图片转成base64编码的dataURL字符串写到代码中
                limit: 10000,
                // 其他的图片转移到 
                name: 'images/[name].[ext]',
                // 如果不改变此项，会导致输出为一个对象
                esModule: false,
            }
        }
    },{
        test: /\.(htm|html)$/,
        loader: 'html-withimg-loader'
    },{
        test: /\.art$/,
        loader: "art-template-loader",
        options: {
            // art-template options (if necessary)
            // @see https://github.com/aui/art-template
        }
    }, {
        test: /\.css$/,
        use: [ MiniCssExtractPlugin.loader, 'css-loader' ],
    },{
        test: /\.(woff2?|eot|ttf|otf)$/,
        loader: 'url-loader',
        options:{
            limit: 10000,
            name: 'fonts/[name].[ext]'
        }
    }      
    ]},
    plugins: [
        new HtmlWebpackPlugin({
            filename:'index.html',
            template:'./src/pages/index/index.art',
            chunks:['index']
        }),
        new HtmlWebpackPlugin({
            filename:'about.html',
            template:'./src/pages/about/about.art',
            chunks:['about']
        }),
        new MiniCssExtractPlugin({
            filename: ({ chunk }) => `css/${chunk.name.replace('/js/', '/css/')}.css`,
        })
    ]
};
