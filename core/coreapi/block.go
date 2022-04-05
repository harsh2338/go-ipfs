package coreapi

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"

	blocks "github.com/ipfs/go-block-format"
	cid "github.com/ipfs/go-cid"
	pin "github.com/ipfs/go-ipfs-pinner"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	caopts "github.com/ipfs/interface-go-ipfs-core/options"
	path "github.com/ipfs/interface-go-ipfs-core/path"

	util "github.com/ipfs/go-ipfs/blocks/blockstoreutil"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockAPI CoreAPI

type BlockStat struct {
	path path.Resolved
	size int
}

func (api *BlockAPI) Put(ctx context.Context, src io.Reader, opts ...caopts.BlockPutOption) (coreiface.BlockStat, error) {
	settings, pref, err := caopts.BlockPutOptions(opts...)
	if err != nil {
		return nil, err
	}
	/**
	--------------------------------------------------
	*/
	conn, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}

	contract, err := NewPlagiarismContract(common.HexToAddress("0x186e74d1FA0c224bd73FA73a384c332Baa10E7f8"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}
	log.Fatalf("Successfully reached")

	// amt, _ := contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0x387fc6939b5e54b2f11793df05388f9d11942948"))
	// BalanceOf(&bind.CallOpts{}, common.HexToAddress("0x387fc6939b5e54b2f11793df05388f9d11942948"))

	var strToConvert string

	strToConvert = "Tutorial Gateway"

	byteString := make([]byte, 16)
	var ret [16]byte
	copy(ret[:], byteString)

	copy(byteString, strToConvert)
	var arrayofarrays [][16]byte
	// for i := range arrayofarrays { //assign
	arrayofarrays[0] = ret
	// }

	copy(byteString, strToConvert)
	res3, _ := contract.UploadFile(&bind.TransactOpts{}, big.NewInt(500), "docCID", "Code file 1", "Testing the contract", "codeFingerPrint", arrayofarrays, big.NewInt(1))

	fmt.Println(res2)
	fmt.Println(res1)
	fmt.Println(res3)

	data, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, err
	}

	bcid, err := pref.Sum(data)
	if err != nil {
		return nil, err
	}

	b, err := blocks.NewBlockWithCid(data, bcid)
	if err != nil {
		return nil, err
	}

	if settings.Pin {
		defer api.blockstore.PinLock(ctx).Unlock(ctx)
	}

	err = api.blocks.AddBlock(ctx, b)
	if err != nil {
		return nil, err
	}

	if settings.Pin {
		api.pinning.PinWithMode(b.Cid(), pin.Recursive)
		if err := api.pinning.Flush(ctx); err != nil {
			return nil, err
		}
	}

	return &BlockStat{path: path.IpldPath(b.Cid()), size: len(data)}, nil
}

func (api *BlockAPI) Get(ctx context.Context, p path.Path) (io.Reader, error) {
	rp, err := api.core().ResolvePath(ctx, p)
	if err != nil {
		return nil, err
	}
	/**
	--------------------------------------------------
	*/
	conn, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}

	contract, err := NewPlagiarismContract(common.HexToAddress("0x186e74d1FA0c224bd73FA73a384c332Baa10E7f8"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	// amt, _ := contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0x387fc6939b5e54b2f11793df05388f9d11942948"))
	// BalanceOf(&bind.CallOpts{}, common.HexToAddress("0x387fc6939b5e54b2f11793df05388f9d11942948"))
	res1, _ := contract.DoesUserHavePermission(&bind.CallOpts{}, "a", common.HexToAddress("a"))
	fmt.Println(res1)

	b, err := api.blocks.GetBlock(ctx, rp.Cid())
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b.RawData()), nil
}

func (api *BlockAPI) Rm(ctx context.Context, p path.Path, opts ...caopts.BlockRmOption) error {
	rp, err := api.core().ResolvePath(ctx, p)
	if err != nil {
		return err
	}

	settings, err := caopts.BlockRmOptions(opts...)
	if err != nil {
		return err
	}
	cids := []cid.Cid{rp.Cid()}
	o := util.RmBlocksOpts{Force: settings.Force}

	out, err := util.RmBlocks(ctx, api.blockstore, api.pinning, cids, o)
	if err != nil {
		return err
	}

	select {
	case res, ok := <-out:
		if !ok {
			return nil
		}

		remBlock, ok := res.(*util.RemovedBlock)
		if !ok {
			return errors.New("got unexpected output from util.RmBlocks")
		}

		if remBlock.Error != "" {
			return errors.New(remBlock.Error)
		}
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (api *BlockAPI) Stat(ctx context.Context, p path.Path) (coreiface.BlockStat, error) {
	rp, err := api.core().ResolvePath(ctx, p)
	if err != nil {
		return nil, err
	}

	b, err := api.blocks.GetBlock(ctx, rp.Cid())
	if err != nil {
		return nil, err
	}

	return &BlockStat{
		path: path.IpldPath(b.Cid()),
		size: len(b.RawData()),
	}, nil
}

func (bs *BlockStat) Size() int {
	return bs.size
}

func (bs *BlockStat) Path() path.Resolved {
	return bs.path
}

func (api *BlockAPI) core() coreiface.CoreAPI {
	return (*CoreAPI)(api)
}
