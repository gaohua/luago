package codegen

import . "luago/compiler/ast"

func cgStat(fi *funcInfo, node Stat) {
	switch stat := node.(type) {
	case *FuncCallStat:
		cgFuncCallStat(fi, stat)
	case *BreakStat:
		cgBreakStat(fi, stat)
	case *DoStat:
		cgDoStat(fi, stat)
	case *RepeatStat:
		cgRepeatStat(fi, stat)
	case *WhileStat:
		cgWhileStat(fi, stat)
	case *IfStat:
		cgIfStat(fi, stat)
	case *ForNumStat:
		cgForNumStat(fi, stat)
	case *ForInStat:
		cgForInStat(fi, stat)
	case *AssignStat:
		cgAssignStat(fi, stat)
	case *LocalVarDeclStat:
		cgLocalVarDeclStat(fi, stat)
	case *LocalFuncDefStat:
		cgLocalFuncDefStat(fi, stat)
	case *LabelStat, *GotoStat:
		panic("not supported!")
	}
}

func cgLocalFuncDefStat(fi *funcInfo, node *LocalFuncDefStat) {
	r := fi.addLocVar(node.Name)
	cgFuncDefExp(fi, node.Exp, r)
}

func cgFuncCallStat(fi *funcInfo, node *FuncCallStat) {
	r := fi.allocReg()
	cgFuncCallExp(fi, node, r, 0)
	fi.freeReg()
}

func cgBreakStat(fi *funcInfo, node *BreakStat) {
	pc := fi.emitJmp(0, 0)
	fi.addBreakJmp(pc)
}

func cgDoStat(fi *funcInfo, node *DoStat) {
	fi.enterScope(false)
	cgBlock(fi, node.Block)
	fi.closeOpenUpvals()
	fi.exitScope()
}

func cgWhileStat(fi *funcInfo, node *WhileStat) {
	pcBeforeExp := fi.pc()
	r := fi.allocReg()
	cgExp(fi, node.Exp, r, 1)
	fi.freeReg()
	fi.emitTest(r, 0)
	pcJmpToEnd := fi.emitJmp(0, 0)
	fi.enterScope(true)
	cgBlock(fi, node.Block)
	fi.closeOpenUpvals()
	fi.emitJmp(0, pcBeforeExp-fi.pc()-1)
	fi.fixSbx(pcJmpToEnd, fi.pc()-pcJmpToEnd)
}

func cgRepeatStat(fi *funcInfo, node *RepeatStat) {
	fi.enterScope(true)

	pcBeforeBlock := fi.pc()
	cgBlock(fi, node.Block)

	r := fi.allocReg()
	cgExp(fi, node.Exp, r, 1)
	fi.freeReg()

	fi.emitTest(r, 0)
	fi.emitJmp(fi.getJmpArgA(), pcBeforeBlock-fi.pc()-1)
	fi.closeOpenUpvals()

	fi.exitScope()
}
