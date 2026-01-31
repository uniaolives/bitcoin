// Copyright (c) 2024-present The Bitcoin Core developers
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

#include <rpc/server.h>
#include <rpc/util.h>
#include <univalue.h>
#include <util/strencodings.h>
#include <script/script.h>
#include <core_io.h>

RPCHelpMan anchorgamestate()
{
    return RPCHelpMan{"anchorgamestate",
        "\nAnchors a game state hash into the Bitcoin blockchain using OP_RETURN.\n",
        {
            {"state_hash", RPCArg::Type::STR_HEX, RPCArg::Optional::NO, "The hex-encoded game state hash (e.g. BLAKE3 hash)."},
        },
        RPCResult{
            RPCResult::Type::OBJ, "", "",
            {
                {RPCResult::Type::STR_HEX, "op_return_script", "The hex-encoded OP_RETURN script containing the game state hash."},
            }
        },
        RPCExamples{
            HelpExampleCli("anchorgamestate", "\"00112233445566778899aabbccddeeff00112233445566778899aabbccddeeff\"")
            + HelpExampleRpc("anchorgamestate", "\"00112233445566778899aabbccddeeff00112233445566778899aabbccddeeff\"")
        },
        [&](const RPCHelpMan& self, const JSONRPCRequest& request) -> UniValue
{
    std::vector<unsigned char> vch_hash = ParseHexV(request.params[0], "state_hash");

    CScript script = CScript() << OP_RETURN << vch_hash;

    UniValue result(UniValue::VOBJ);
    result.pushKV("op_return_script", HexStr(script));

    return result;
},
    };
}

void RegisterGameRPCCommands(CRPCTable &t)
{
    static const CRPCCommand commands[]{
        {"blockchain", anchorgamestate},
    };
    for (const auto& c : commands) {
        t.appendCommand(c.name, &c);
    }
}
