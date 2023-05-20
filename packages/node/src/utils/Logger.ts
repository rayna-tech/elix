import chalk from "chalk"

export default (msg: string) => {
    const p = " Elix "
    return {
        info: () => {
            console.log(`${chalk.bgCyanBright(p)} ${msg}`)
        },
        error: () => {
            console.log(`${chalk.bgRedBright(p)} ${msg}`)
        },
        rest: () => {
            console.log(`${chalk.bgMagentaBright(p)} ${msg}`)
        },
        warn: () => {
            console.log(`${chalk.bgYellowBright(p)} ${msg}`)
        }
    }
}