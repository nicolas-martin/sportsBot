'use strict';
let exec = require('child_process').exec;

exports.handler = (event, context) => {
    const child = exec("./go_compiled_code " + event.userid, (error) => {
        if(error){
            console.log('Error:',error);
            context.fail(error);
        }
    });
    child.stdout.on('data',function(data){
        console.log('Data:',data);
        context.succeed(data)
    });
};
