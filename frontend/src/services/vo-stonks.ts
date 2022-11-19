/* eslint:disable */
// Code generated by gotsrpc https://github.com/foomo/gotsrpc/v2  - DO NOT EDIT.
import * as github_com_hackaTUM_GameOfStonks_services_stonks from './vo-stonks'; // ./frontend/src/services/vo-stonks.ts to ./frontend/src/services/vo-stonks.ts
// github.com/hackaTUM/GameOfStonks/services/stonks.DataPoint
export interface DataPoint {
	Time:number;
	Value:number;
}
// github.com/hackaTUM/GameOfStonks/services/stonks.DataPoints
export type DataPoints = Array<github_com_hackaTUM_GameOfStonks_services_stonks.DataPoint>
// github.com/hackaTUM/GameOfStonks/services/stonks.Err
export interface Err {
	message:string;
}
// github.com/hackaTUM/GameOfStonks/services/stonks.Match
export interface Match {
	UserSell:string;
	UserBuy:string;
	Quantity:number;
	TimeStamp:number;
}
// github.com/hackaTUM/GameOfStonks/services/stonks.Order
export interface Order {
	UserName:string;
	OrderType:github_com_hackaTUM_GameOfStonks_services_stonks.OrderType;
	Quantity:number;
	TimeStamp:number;
}
// github.com/hackaTUM/GameOfStonks/services/stonks.OrderType
export enum OrderType {
	Buy = "buy",
	Sell = "sell",
}
// github.com/hackaTUM/GameOfStonks/services/stonks.PlaceOrderCmd
export interface PlaceOrderCmd {
	Stonk:github_com_hackaTUM_GameOfStonks_services_stonks.StonkName;
	Quantity:number;
	Price:number;
	OrderType:github_com_hackaTUM_GameOfStonks_services_stonks.OrderType;
}
// github.com/hackaTUM/GameOfStonks/services/stonks.StonkInfo
export interface StonkInfo {
	Name:github_com_hackaTUM_GameOfStonks_services_stonks.StonkName;
	TimeSeries:Array<github_com_hackaTUM_GameOfStonks_services_stonks.DataPoint>|null;
	MatchHistory:Array<github_com_hackaTUM_GameOfStonks_services_stonks.Match>|null;
	UserOrders:Array<github_com_hackaTUM_GameOfStonks_services_stonks.Order>|null;
	Orders:Array<github_com_hackaTUM_GameOfStonks_services_stonks.Order>|null;
}
// github.com/hackaTUM/GameOfStonks/services/stonks.StonkName
export enum StonkName {
	StonkHouse = "house",
	StonkMate = "mate",
	StonkPaperClip = "paperClip",
	StonkPencil = "pencil",
	StonkScissors = "scissors",
	StonkEmpty = "",
}
// github.com/hackaTUM/GameOfStonks/services/stonks.UpdateOrderCmd
export interface UpdateOrderCmd {
	Id:string;
	Quantity:number;
	Price:number;
}
// github.com/hackaTUM/GameOfStonks/services/stonks.User
export interface User {
	Name:string;
	Money:number;
	ReservedMoney:number;
	Stonks:Record<github_com_hackaTUM_GameOfStonks_services_stonks.StonkName,number>|null;
	NetWorth:number;
	NetWorthTimeSeries:github_com_hackaTUM_GameOfStonks_services_stonks.DataPoints|null;
}
// end of common js